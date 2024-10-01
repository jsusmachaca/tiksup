from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType, ArrayType
from pyspark.sql import functions as F
from os import environ
from time import sleep
import uuid

class SparkProcess:
    def __init__(self) -> None:
        while True:
            try:
                self.spark = SparkSession.builder \
                    .appName("MyApp") \
                    .master(f"spark://{environ.get("SPARK_HOST")}:{environ.get("SPARK_PORT")}") \
                    .getOrCreate()
                break
            except Exception:
                print("Esperando a que Spark esté disponible...")
                sleep(1)

        self.movie_schema = StructType([
            StructField("id", StringType(), True),
            StructField("url", StringType(), True),
            StructField("genre", ArrayType(StringType()), True),
            StructField("director", StringType(), True),
            StructField("protagonist", StringType(), True),
            StructField("title", StringType(), True)
        ])

    def recommend_movies(self, user_preferences, movies):
        movies_df = self.spark.createDataFrame(movies, schema=self.movie_schema)

        genre_scores = user_preferences.get("preferences", {}).get("genre_score", [])
        director_scores = user_preferences.get("preferences", {}).get("director_score", [])
        protagonist_scores = user_preferences.get("preferences", {}).get("protagonist_score", [])

        genre_scores_df = self.spark.createDataFrame(
            [{"name": item["name"], "score": f"{item['score']:.2f}"} for item in genre_scores]
        ).withColumnRenamed("score", "genre_score")

        director_scores_df = self.spark.createDataFrame(
            [{"name": item["name"], "score": f"{item['score']:.2f}"} for item in director_scores]
        ).withColumnRenamed("score", "director_score")

        protagonist_scores_df = self.spark.createDataFrame(
            [{"name": item["name"], "score": f"{item['score']:.2f}"} for item in protagonist_scores]
        ).withColumnRenamed("score", "protagonist_score")

        recommendations_df = movies_df \
            .join(genre_scores_df, F.array_contains(movies_df.genre, genre_scores_df.name), "left") \
            .join(director_scores_df, movies_df.director.contains(director_scores_df.name), "left") \
            .join(protagonist_scores_df, movies_df.protagonist.contains(protagonist_scores_df.name), "left")

        recommendations_df = recommendations_df.withColumn(
            "combined_score",
            0.4 * F.coalesce(recommendations_df.genre_score, F.lit(0)) +
            0.3 * F.coalesce(recommendations_df.director_score, F.lit(0)) +
            0.3 * F.coalesce(recommendations_df.protagonist_score, F.lit(0))
        )

        sorted_recommendations_df = recommendations_df.filter(recommendations_df.combined_score > 0) \
            .orderBy(F.desc("combined_score"))

        sorted_recommendations = sorted_recommendations_df.collect()

        result = {
            "user_id": user_preferences["user_id"],
            "movies": [{
                "id": row.id,
                "url": row.url,
                "title": row.title,
                "genre": row.genre,
                "director": row.director,
                "protagonist": row.protagonist,
            } for row in sorted_recommendations]
        }

        return result
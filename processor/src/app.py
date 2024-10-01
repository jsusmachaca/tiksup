import json
import redis
import time
import uuid
from fastapi import FastAPI
from pyspark.sql import SparkSession
from pyspark.sql.types import StructType, StructField, StringType
from os import environ
from dotenv import load_dotenv

load_dotenv()

app = FastAPI()

redis_client = redis.StrictRedis(host=environ.get("REDIS_HOST"), port=environ.get("REDIS_PORT"), db=0)

while True:
    try:
        spark = SparkSession.builder \
            .appName("MyApp") \
            .master(f"spark://{environ.get("SPARK_HOST")}:{environ.get("SPARK_PORT")}") \
            .getOrCreate()
        break
    except Exception as e:
        print("Esperando a que Spark esté disponible...")
        time.sleep(5)

movie_schema = StructType([
    StructField("movie_id", StringType(), True),
    StructField("url", StringType(), True),
    StructField("genre", StringType(), True),
    StructField("director", StringType(), True),
    StructField("protagonist", StringType(), True),
    StructField("title", StringType(), True)
])

def recommend_movies(user_preferences, movies):
    movies_df = spark.createDataFrame(movies, schema=movie_schema)

    genre_scores = user_preferences.get("preferences", {}).get("genre_score", [])
    director_scores = user_preferences.get("preferences", {}).get("director_score", [])
    protagonist_scores = user_preferences.get("preferences", {}).get("protagonist_score", [])

    genre_scores_df = spark.createDataFrame(genre_scores).withColumnRenamed("score", "genre_score")
    director_scores_df = spark.createDataFrame(director_scores).withColumnRenamed("score", "director_score")
    protagonist_scores_df = spark.createDataFrame(protagonist_scores).withColumnRenamed("score", "protagonist_score")

    recommendations_df = movies_df.join(genre_scores_df, movies_df.genre.contains(genre_scores_df.name), "left") \
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
            "movie_id": row.movie_id if row.movie_id is not None else str(uuid.uuid4()),  
            "url": row.url,
            "title": row.title,
            "genre": row.genre,
            "director": row.director,
            "protagonist": row.protagonist,
            "score": row.combined_score
        } for row in sorted_recommendations]
    }

    return result

@app.post("/recommend")
async def receive_data(data: dict):
    user_preferences = {
        "user_id": str(data.get("user_id")),
        "preferences": {
            "genre_score": data.get("preferences", {}).get("genre_score", []),
            "protagonist_score": data.get("preferences", {}).get("protagonist_score", []),
            "director_score": data.get("preferences", {}).get("director_score", []),
        }
    }

    movies = data.get("movies", [])
    recommendations = recommend_movies(user_preferences, movies)

    redis_client.set(f"user:{user_preferences['user_id']}:recommendations", json.dumps(recommendations))

    return {"message": "Recived data"}

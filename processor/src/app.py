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
    recommendations = []
    non_recommended = []

    for movie in movies:
        movie_genre = ", ".join(movie["genre"])  
        movie_director = movie["director"]
        movie_protagonist = movie["protagonist"]
        movie_title = movie["title"]

        genre_scores = user_preferences.get("preferences", {}).get("genre_score", [])
        genre_score = next((score["score"] for score in genre_scores if score["name"].lower() in movie_genre.lower()), 0)

        director_scores = user_preferences.get("preferences", {}).get("director_score", [])
        director_score = next((score["score"] for score in director_scores if score["name"].lower() in movie_director.lower()), 0)

        protagonist_scores = user_preferences.get("preferences", {}).get("protagonist_score", [])
        protagonist_score = next((score["score"] for score in protagonist_scores if score["name"].lower() in movie_protagonist.lower()), 0)

        combined_score = 0.4 * genre_score + 0.3 * director_score + 0.3 * protagonist_score

        movie_entry = {
            "movie_id": movie.get("id", str(uuid.uuid4())),  
            "url": movie["url"],
            "title": movie_title,
            "genre": movie_genre,
            "director": movie_director,
            "protagonist": movie_protagonist,
            "score": combined_score
        }

        if combined_score > 0:
            recommendations.append(movie_entry)
        else:
            non_recommended.append(movie_entry)

    sorted_recommendations = sorted(recommendations, key=lambda x: x["score"], reverse=True)
    all_movies_ordered = sorted_recommendations + non_recommended

    result = {
        "user_id": user_preferences["user_id"],
        "movies": all_movies_ordered
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

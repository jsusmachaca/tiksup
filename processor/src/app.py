import json
import redis
from fastapi import FastAPI
from os import environ
from dotenv import load_dotenv
from .service import SparkProcess

load_dotenv()

app = FastAPI()
spark = SparkProcess()

redis_client = redis.StrictRedis(host=environ.get("REDIS_HOST"), port=environ.get("REDIS_PORT"), db=0)

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

    recommendations = spark.recommend_movies(user_preferences, movies)

    redis_client.set(f"user:{user_preferences['user_id']}:recommendations", json.dumps(recommendations))

    return {"message": "Recived data"}

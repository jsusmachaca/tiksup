import json
import redis
from fastapi import FastAPI, BackgroundTasks
from os import environ
from dotenv import load_dotenv
from .service import SparkProcess

load_dotenv()

app = FastAPI()

spark = SparkProcess()
redis_client = redis.StrictRedis(
    host=environ.get("REDIS_HOST"),
    port=environ.get("REDIS_PORT"),
    db=0
)

@app.get("/hello-world")
async def hello_world():
    return {"message": "hello world"}

def process_recommendation(user_preferences, movies) -> bool:
    recommendations = spark.recommend_movies(user_preferences, movies)

    if recommendations['user_id'] is None:
        print("Warning: user_id is None, not storing in Redis")
        return

    result = redis_client.set(f"user:{recommendations['user_id']}:recommendations", json.dumps(recommendations))
    
    if result:
        print("Stored in Redis")
        return True
    return False
    
@app.post("/recommend")
async def receive_data(data: dict, background_tasks: BackgroundTasks):
    user_id = str(data.get("user_id"))
    if not user_id:
        return {"error": "user_id is required."}

    user_preferences = {
        "user_id": user_id,
        "preferences": {
            "genre_score": data.get("preferences", {}).get("genre_score", []),
            "protagonist_score": data.get("preferences", {}).get("protagonist_score", []),
            "director_score": data.get("preferences", {}).get("director_score", []),
        }
    }

    movies = data.get("movies", [])

    background_tasks.add_task(process_recommendation, user_preferences, movies)

    return {"message": "Received data, processing in background."}

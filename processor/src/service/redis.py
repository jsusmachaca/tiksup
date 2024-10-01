

def recommend_movies(user_preferences, movies):
    recommendations = []
    non_recommended = []

    for movie in movies:
        movie_id = movie.get("id")
        movie_url = movie["url"]
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
            "movie_id": movie_id,  
            "url": movie_url,
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
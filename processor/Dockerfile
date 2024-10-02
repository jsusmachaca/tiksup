FROM python:3.12-slim

RUN apt-get update && apt-get install -y \
    default-jre-headless \
    procps \
    && rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME=/usr/lib/jvm/default-java
ENV PATH="$JAVA_HOME/bin:$PATH"

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY src src

CMD ["uvicorn", "src.app:app", "--host", "0.0.0.0", "--port", "8000"]

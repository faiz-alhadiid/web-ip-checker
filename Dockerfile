FROM python:3.8-slim
RUN apt-get update && apt-get install -y dnsutils
COPY . /app
WORKDIR /app
RUN pip install --no-cache-dir -r requirements.txt
CMD exec gunicorn --bind :$PORT --workers 1 --threads 1 --timeout 0 main:app
FROM python:3.12-alpine

# Устанавливаем необходимые системные зависимости
RUN apk add --no-cache gcc musl-dev zlib-dev libffi-dev openssl-dev

WORKDIR /src

COPY requirements.txt requirements.txt

RUN pip install --upgrade pip && \
    pip install -r requirements.txt

COPY . .

CMD python -m main
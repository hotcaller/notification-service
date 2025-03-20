FROM python:3.13-slim

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файл с зависимостями в контейнер
COPY requirements.txt .

# Устанавливаем зависимости
RUN pip install --no-cache-dir -r requirements.txt

# Копируем всё содержимое директории server в контейнер
COPY . .

# Добавляем /app в PYTHONPATH
ENV PYTHONPATH="${PYTHONPATH}:/app"

# Открываем порт 8000
EXPOSE 8000

# Команда для запуска приложения
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]

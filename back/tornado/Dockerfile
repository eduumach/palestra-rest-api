# Use an official Python runtime as a parent image
FROM python:3.9

# Set the working directory in the container
WORKDIR /app

# Copy your Tornado application files to the container
COPY . /app

# Install any necessary dependencies (requirements.txt) for your Tornado application
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Expose the port on which your Tornado application will run
EXPOSE 8888

# Define the command to start your Tornado application
CMD ["python", "main.py"]

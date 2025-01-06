import os

class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'default_secret_key'
    SQLALCHEMY_DATABASE_URI = r'D:\SanskrutiKakad_nexturn_training _program\M3_Python_and_API_assignments\1_Python_API_Exercise\1.1_BookBuddy\mydb.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False
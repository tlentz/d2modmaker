import csv
import os

class D2Data:
    fileData = []

    def __init__(self):
        self.fileData = []
        uniques = D2File("UniqueItems.txt")
        self.fileData.append(uniques)

class D2File:
    name = "sdfasdfsf"
    headers = {}
    rows = []

    def __init__(self, name):
        self.name = name

    def read(self):
        thisDir = os.path.dirname(os.path.realpath('__file__'))
        fileDir = os.path.join(thisDir, f"../assets/113c-data/{self.name}")
        print(fileDir)

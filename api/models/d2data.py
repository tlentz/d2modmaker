import csv
import os
import json
import pandas

def get_data_dir():
    this_dir = os.path.dirname(os.path.realpath('__file__'))
    return os.path.join(this_dir, "../api/assets/d2-src/")

class D2Data:
    def __init__(self):
        self.files = self.read_files()

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, indent=4)

    def read_files(self):
        files = {}
        for f in os.listdir(get_data_dir()):
            files[f] = D2File(f)
        return files


class D2File:
    def __init__(self, name):
        self.name = name
        self.file_name = f"{get_data_dir()}{self.name}"

        self.fieldnames = []
        self.data = []

        self.read()

    def read(self):
        self.data = []
        print (self.name)
        with open(self.file_name, newline='', encoding="Windows-1252") as csvfile:
            reader = csv.DictReader(f=csvfile, delimiter='\t')
            self.fieldnames = reader.fieldnames
            for row in reader:
                self.data.append(row)

    def write(self):
        with open(self.file_name, 'w', newline='', encoding="Windows-1252") as csvfile:
           writer = csv.DictWriter(csvfile, fieldnames=self.fieldnames, delimiter='\t')
           writer.writeheader()
           writer.writerows(self.data)

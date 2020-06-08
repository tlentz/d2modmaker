import csv
import pandas as pd

src_data = {}
src_files = [
    "Gems",
    "Runes",
    "SetItems",
    "Sets",
    "UniqueItems"
]


def read_src_file(fname):
    csvfile = f"./source/{fname}.txt"
    df = pd.read_csv(csvfile, delimiter="\t", engine="python")
    src_data[fname] = df


def main():
    for f in src_files:
        read_src_file(f)
    for key in src_data:
        print(key)
        print(src_data[key].columns)


if __name__ == "__main__":
    # execute only if run as a script
    main()

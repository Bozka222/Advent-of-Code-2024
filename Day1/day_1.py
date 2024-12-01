import csv

FILE_PATH = "input_data.csv"
column1 = []
column2 = []

with open(FILE_PATH, "r", encoding="utf-8") as input_data:
    csvreader = csv.reader(input_data)

    for row in csvreader:
        column1.append(int(row[0]))
        column2.append(int(row[1]))

column1.sort()
column2.sort()

# print(column1)
# print(column2)

total_distance = sum(abs(column1[i] - column2[i]) for i in range(len(column1)))
print(total_distance)

import csv

file_path = "input_data.csv"
column1 = []
column2 = []

with open(file_path, "r") as input_data:
    csvreader = csv.reader(input_data)

    for row in csvreader:
        column1.append(int(row[0]))
        column2.append(int(row[1]))

count = 0
total_distance = 0
for number in column1:
    for number2 in column2:
        if number == number2:
            count += 1

    total_distance += (number * count)
    count = 0

print(total_distance)

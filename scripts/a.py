import matplotlib.pyplot as plt

xs = list(range(1,18))
ys = [
    1,
    2,
    5,
    7,
    12,
    22,
    45,
    90,
    180,
    360,
    720,
    1066,
    1519,
    2170,
    2941,
    3957,
    5170
]


def ml(x):
    return int(x * 1.10)

yss = []

for idx, val in enumerate(ys):
    if idx > 10:
        yss.append(ml(ys[idx - 1]))
    else:
        yss.append(val)


print(yss)
# plt.plot(xs, new_ys)
# plt.xlabel('Runes x')
# plt.ylabel('Rune Drop Chance')
# plt.show()
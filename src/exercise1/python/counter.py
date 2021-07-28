from threading import Thread

counter = 0
rounds = 100000


def run():
    for i in range(rounds):
        global counter
        counter += 1


if __name__ == '__main__':
    t1 = Thread(target=run)
    t2 = Thread(target=run)

    t1.start()
    t2.start()

    t1.join()
    t2.join()
    print(counter)


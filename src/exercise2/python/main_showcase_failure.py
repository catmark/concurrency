from backend import Backend
from month import Month
from threading import Thread
import time

if __name__ == '__main__':
    backend = Backend()

    threads = []

    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.July)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.July)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.July)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.July)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.August)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.August)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.August)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.August)))
    threads.append(Thread(target= backend.on_ad_click, args=(1,Month.August)))

    for thread in threads:
        thread.start()
        time.sleep(0.03)


    for thread in threads:
        thread.join()

    backend.cash_out(1, Month.July)
    backend.cash_out(1, Month.August)


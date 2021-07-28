from backend import Backend
from month import Month

if __name__ == '__main__':
    backend = Backend()

    backend.on_ad_click(1,Month.July)
    backend.on_ad_click(1,Month.July)
    backend.on_ad_click(1,Month.July)
    backend.on_ad_click(1,Month.July)
    backend.on_ad_click(1,Month.August)
    backend.on_ad_click(1,Month.August)
    backend.on_ad_click(1,Month.August)
    backend.on_ad_click(1,Month.August)
    backend.on_ad_click(1,Month.August)

    backend.cash_out(1,Month.July)
    backend.cash_out(1,Month.August)


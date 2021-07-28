import time
from month import Month


def initTransactionsForUser():
    return {
        Month.January: [],
        Month.February: [],
        Month.March: [],
        Month.April: [],
        Month.May: [],
        Month.June: [],
        Month.July: [],
        Month.August: [],
        Month.September: [],
        Month.October: [],
        Month.November: [],
        Month.December: []
    }


class DatabaseMock:
    balance = {
        1: 0.00,#user 1
        2: 0.00,#user 2
        3: 0.00,#user 3
        4: 0.00,#user 4
    }

    transactions = {
        1: initTransactionsForUser(),#user 1
        2: initTransactionsForUser(),#user 2
        3: initTransactionsForUser(),#user 3
        4: initTransactionsForUser(),#user 4
    }

    @staticmethod
    def getUserBalance(userId: int):
        balance = DatabaseMock.balance[userId]
        time.sleep(0.1)
        return balance

    @staticmethod
    def getUserTransactionsForMonth(userId: int, month: Month):
        return DatabaseMock.transactions[userId][month]

    @staticmethod
    def setUserBalance(userId: int, balance: float):
        DatabaseMock.balance[userId] = balance

    @staticmethod
    def addUserTransactionForMonth(userId: int, month: Month, transaction: float):
        DatabaseMock.transactions[userId][month].append(transaction)

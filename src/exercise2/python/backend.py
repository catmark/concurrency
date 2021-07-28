from databaseMock import DatabaseMock
from month import Month


class Backend:
    def on_ad_click(self, userId: int, currentMonth: Month):
        valueOfSingleClick = 0.001
        DatabaseMock.addUserTransactionForMonth(userId, currentMonth, valueOfSingleClick)

        # recalculate balance
        balance = DatabaseMock.getUserBalance(userId)
        balance += valueOfSingleClick
        DatabaseMock.setUserBalance(userId, balance)

    def cash_out(self, userId: int, month: Month):
        balance = DatabaseMock.getUserBalance(userId)
        if balance > 0:
            #calculate amount to be paid to user
            transactions = DatabaseMock.getUserTransactionsForMonth(userId, month)
            amountToBePaid = sum(transactions)

            self._make_payment(month, userId, amountToBePaid)

            #recalculate balance
            balance = DatabaseMock.getUserBalance(userId)
            balance -= amountToBePaid
            DatabaseMock.setUserBalance(userId, balance)

    def _make_payment(self, month: Month, userId: int, amountToBePaid: float):
        # some logic to make actual payment
        print(f"Payment made ({month}): userId={userId}, amount={amountToBePaid}, balance_before_payment={DatabaseMock.getUserBalance(userId):.3f}, balance_after_payment={DatabaseMock.getUserBalance(userId)-amountToBePaid:.3f}")

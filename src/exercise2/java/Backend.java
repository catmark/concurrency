import java.time.Month;
import java.util.List;

class Backend {
    private Database db = new Database();

    void onAddClick(int userId, Month month) {
        double valueOfSingleClick = 0.001;
        db.addUserTransaction(userId, month, valueOfSingleClick);

        //recalculate balance
        double balance = db.getUserBalance(userId);
        balance += valueOfSingleClick;
        db.setUserBalance(userId, balance);
    }

    void cashOut(int userId, Month month) {
        double balance = db.getUserBalance(userId);
        if (balance > 0) {
            List<Double> transactions = db.getUserTransactions(userId, month);
            Double amountToBePaid = transactions.stream().reduce(0.0, Double::sum);

            this.makePayment(month, userId, amountToBePaid);

            //recalculate balance
            balance = db.getUserBalance(userId);
            balance -= amountToBePaid;
            db.setUserBalance(userId, balance);
        }
    }

    private void makePayment(Month month, int userId, Double amount) {
        Double balance = db.getUserBalance(userId);
        System.out.printf(
                "Payment made (%.3s): userId=%d, amount=%.3f, balance_before_payment=%.3f, balance_after_payment=%.3f\n",
                month, userId, amount, balance, balance - amount);
    }
}

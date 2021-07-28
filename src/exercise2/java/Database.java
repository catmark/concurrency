import java.time.Month;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Database {
    private Map<Integer, Double> balance = new HashMap<>();
    private Map<Integer, Map<Month, List<Double>>> transactions = new HashMap<>();


    Database() {
        balance.put(1, 0.0);
        balance.put(2, 0.0);

        transactions.put(1, emptyTransactions());
        transactions.put(2, emptyTransactions());
    }

    double getUserBalance(int userId) {
        double balance = this.balance.get(userId);
        try {
            Thread.sleep(100);
        } catch (InterruptedException e) {}
        return balance;
    }

    void setUserBalance(int userId, double balance) {
        this.balance.put(userId, balance);
    }

    List<Double> getUserTransactions(int userId, Month month) {
        return transactions.get(userId).get(month);
    }

    void addUserTransaction(int userId, Month month, double transaction) {
        transactions.get(userId).get(month).add(transaction);
    }

    private Map<Month, List<Double>> emptyTransactions() {
        Map<Month, List<Double>> transactions = new HashMap<>();
        for (Month month : Month.values()) {
            transactions.put(month, new ArrayList<>());
        }
        return transactions;
    }
}

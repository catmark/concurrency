import java.time.Month;

public class Swindle {

    public static void main(String[] args) {
        Backend backend = new Backend();
        int userId = 1;

        backend.onAddClick(userId, Month.JULY);
        backend.onAddClick(userId, Month.JULY);
        backend.onAddClick(userId, Month.JULY);
        backend.onAddClick(userId, Month.JULY);
        backend.onAddClick(userId, Month.JULY);
        backend.onAddClick(userId, Month.AUGUST);
        backend.onAddClick(userId, Month.AUGUST);
        backend.onAddClick(userId, Month.AUGUST);
        backend.onAddClick(userId, Month.AUGUST);

        backend.cashOut(userId, Month.JULY);
        backend.cashOut(userId, Month.AUGUST);
    }
}

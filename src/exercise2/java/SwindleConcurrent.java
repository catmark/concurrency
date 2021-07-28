import java.time.Month;
import java.util.List;

public class SwindleConcurrent {
    public static void main(String[] args) {
        Backend backend = new Backend();
        int userId = 1;

        List<Thread> threads = List.of(
            new Thread(() -> backend.onAddClick(userId, Month.JULY)),
            new Thread(() -> backend.onAddClick(userId, Month.JULY)),
            new Thread(() -> backend.onAddClick(userId, Month.JULY)),
            new Thread(() -> backend.onAddClick(userId, Month.JULY)),
            new Thread(() -> backend.onAddClick(userId, Month.JULY)),
            new Thread(() -> backend.onAddClick(userId, Month.AUGUST)),
            new Thread(() -> backend.onAddClick(userId, Month.AUGUST)),
            new Thread(() -> backend.onAddClick(userId, Month.AUGUST)),
            new Thread(() -> backend.onAddClick(userId, Month.AUGUST))
        );
        threads.forEach((thread) -> {
            thread.start();
            try {
                Thread.sleep(30);
            } catch (InterruptedException e) { }
        });

        threads.forEach((thread) -> {
            try {
                thread.join();
            } catch (InterruptedException e) { }
        });

        backend.cashOut(userId, Month.JULY);
        backend.cashOut(userId, Month.AUGUST);
    }
}

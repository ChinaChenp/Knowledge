import java.util.Arrays;

public class TauTog {
    public static void main(String[] args) {
        int arr[] = new int[5];

        Arrays.fill(arr, 100);

        for (int i = 0; i < arr.length; i++) {
            System.out.println(arr[i]);
        }

        // sort

        int sorts[] = new int[] {5, 9, 1, 4, 2, 3, 10};
        Arrays.sort(sorts);
        for (int i = 0; i < arr.length; i++) {
            System.out.printf("%d, ", sorts[i]);
        }
        System.out.println();
    }
}
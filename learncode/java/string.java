import java.util.Date;
public class StringTest {
    public static void main(String[] args) {
        String str = "fjdaljfda";
        System.out.println(str);

        int len = str.length();
        System.out.println(len);

        String str1 = new String("chenping is good");
        System.out.println(str1);

        System.out.println(str1.lastIndexOf("good"));

        // 分割
        String str2 = new String("a, b, c, d, e, f,g");
        String[] re = str2.split(",", 4);
        for (int i = 0; i < re.length; i++) {
            System.out.print(re[i]);
        }
        System.out.println();
        
        Date date = new Date();
        String date_for = String.format("%tY-%tB-%td", date, date, date);
        String date_for1 = String.format("%4d-%02d-%02d", date.getYear()+1900, date.getMonth(), date.getDay());
        System.out.println(date_for);
        System.out.println(date_for1);
    }
}
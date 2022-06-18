using System.Linq;
public class Solution {
    public int CountPrimes(int n) {
        if (n <= 2 )
        {
            return 0;
        }

        List<int> currentPrimes = new List<int> {2};
        for (int x = 3; x < n; x+=2)
        {
            if (currentPrimes.Any(c => x % c == 0))
                continue;
            currentPrimes.Add(x);
        }
        return currentPrimes.Count;
    }
}
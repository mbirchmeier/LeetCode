﻿using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

class Result
{

    /*
     * Complete the 'plusMinus' function below.
     *
     * The function accepts INTEGER_ARRAY arr as parameter.
     */

    public static void plusMinus(List<int> arr)
    {
        float neg = 0;
        float zero = 0;
        float pos = 0;
        foreach (int i in arr)
        {
            if (i == 0)
            {
                zero++;
            }
            else if (i < 0)
            {
                neg++;
            }
            else
            {
                pos++;
            }
        }
        float total = arr.Count;

        neg /= total;
        pos /= total;
        zero /= total;

        Console.WriteLine(pos.ToString("0.000000"));
        Console.WriteLine(neg.ToString("0.000000"));
        Console.WriteLine(zero.ToString("0.000000"));

    }

    public static void miniMaxSum(List<int> arr)
    {
        long sum = arr.Sum(x => (long)x);
        long maxSum = sum - arr.Min();
        long minSum = sum - arr.Max();

        string builtString = minSum + " " + maxSum; //use a stringbuilder in prod code to ensure good performance

        Console.WriteLine(builtString);

    }

    public static string timeConversion(string s)
    {
        TimeOnly inTime = TimeOnly.ParseExact(s, "h:mm:ss tt");
        return inTime.ToString("HH:mm:ss");
    }

    public static int lonelyinteger(List<int> a)
    {
        var grouped = a.GroupBy(x => x);
        foreach (var group in grouped)
        {
            if (group.Count() == 1)
            {
                return group.Key;
            }
        }

        return -1;
    }

    public static int getUniqueCharacter(string s)
    {
        int[] hash = new int[26];
        for (int c = 0; c != 26; c++)
        {
            hash[c] = 0;
        }
        for (int x = 0; x != s.Length; x++)
        {
            int current_hash = (int)s[x] - (int)'a';
            hash[current_hash]++;
        }

        int toReturn = -1;

        for (int c = 0; c != 26; c++)
        {
            if (hash[c] == 1)
            {
                char finder = (char)(c + (int)'a');
                int currentIndex = s.IndexOf(finder);
                if (toReturn == -1 || currentIndex < toReturn)
                {
                    toReturn = currentIndex;
                }
            }
        }
        return toReturn;
    }

    public static void Euler1()
    {
        int t = Convert.ToInt32(Console.ReadLine());
        for (int a0 = 0; a0 < t; a0++)
        {
            int n = Convert.ToInt32(Console.ReadLine());

            int n = Convert.ToInt32(Console.ReadLine());
            
            n = n - 1;
            
            long threes = n/3;
            threes = (threes * (threes + 1)) / 2 * 3;
            
            long fives = n/5;
            fives = (fives * (fives + 1)) / 2 * 5;
            
            long fifteens = n/15;
            fifteens = (fifteens * (fifteens + 1)) / 2 * 15;
            Console.WriteLine(threes + fives - fifteens);
        }
    }
}

class Solution
{
    public static void Main(string[] args)
    {
        Result.Euler1();
    }
}

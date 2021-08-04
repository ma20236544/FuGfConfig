using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.IO;
using System.Linq;
using System.Text;

namespace ConfGenerate
{
    public class Program
    {
        public static void Main(string[] args)
        {
            const string inFilePath = @"F:\CodeFile\Project\FuGfConfig\ConfGenerate\ConfGenerate\in.txt";
            const string outFilePath = @"F:\CodeFile\Project\FuGfConfig\ConfGenerate\ConfGenerate\out.txt";

            string line;

            // 按行读取
            StreamReader inFile;
            StreamWriter outFile;
            try
            {
                inFile = new StreamReader(inFilePath);
            }
            catch (Exception e)
            {
                Console.WriteLine(e);
                throw;
            }

            var data = new SortedSet<string>();
            while ((line = inFile.ReadLine()) != null)
            {
                var _l = line.Replace(" ", ""); //去掉空格
                var rule = new StringBuilder();
                rule.Append("DOMAIN-SUFFIX,");
                rule.Append(_l);
                rule.Append(",REJECT");
                data.Add(rule.ToString());
            }

            inFile.Close();

            outFile = new StreamWriter(outFilePath);
            foreach (var i in data)
            {
                outFile.WriteLine(i);
            }

            outFile.Close();

            Console.WriteLine("完成");
        }
    }
}
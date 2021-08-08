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

            #region init

            // string line;
            //
            // // 按行读取
            // StreamReader inFile;
            // StreamWriter outFile;
            // try
            // {
            //     inFile = new StreamReader(inFilePath);
            // }
            // catch (Exception e)
            // {
            //     Console.WriteLine(e);
            //     throw;
            // }
            //
            // var data = new SortedSet<string>();
            // while ((line = inFile.ReadLine()) != null)
            // {
            //     var _l = line.Replace(" ", ""); //去掉空格
            //     var rule = new StringBuilder();
            //     rule.Append("DOMAIN-SUFFIX,");
            //     rule.Append(_l);
            //     rule.Append(",REJECT");
            //     data.Add(rule.ToString());
            // }
            //
            // inFile.Close();
            //
            // outFile = new StreamWriter(outFilePath);
            // foreach (var i in data)
            // {
            //     outFile.WriteLine(i);
            // }
            //
            // outFile.Close();

            #endregion


            const string initFilePath = @"F:\CodeFile\Project\FuGfConfig\ConfGenerate\ConfGenerate\init.txt";
            const string in1FilePath = @"F:\CodeFile\Project\FuGfConfig\ConfGenerate\ConfGenerate\in1.txt";
            const string in2FilePath = @"F:\CodeFile\Project\FuGfConfig\ConfGenerate\ConfGenerate\in2.txt";
            const string in3FilePath = @"F:\CodeFile\Project\FuGfConfig\ConfGenerate\ConfGenerate\in3.txt";

            var initReader = new StreamReader(initFilePath);
            var in1Reader = new StreamReader(in1FilePath);
            var in2Reader = new StreamReader(in2FilePath);
            var in3Reader = new StreamReader(in3FilePath);


            var data = new SortedSet<string>();
            var ans = new SortedSet<string>();
            string line;
            while ((line = initReader.ReadLine()) != null)
            {
                var s = line.Remove(line.LastIndexOf(','));
                data.Add(s);
            }

            initReader.Close();
            line = string.Empty;
            // while ((line = in1Reader.ReadLine()) != null)
            // {
            //     var s = line.Remove(line.LastIndexOf(','));
            //     if (!data.Contains(s))
            //     {
            //         var _l = new StringBuilder();
            //         _l.Append(s);
            //         _l.Append(",CustomAdRules");
            //         ans.Add(_l.ToString());
            //     }
            // }
            //
            // in1Reader.Close();
            //
            // while ((line = in2Reader.ReadLine()) != null)
            // {
            //     var s = line.Remove(line.LastIndexOf(','));
            //     if (!data.Contains(s))
            //     {
            //         var _l = new StringBuilder();
            //         _l.Append(s);
            //         _l.Append(",CustomAdRules");
            //         ans.Add(_l.ToString());
            //     }
            // }
            //
            // in2Reader.Close();

            while ((line = in3Reader.ReadLine()) != null)
            {
                var s = line;
                if (!data.Contains(s))
                {
                    var _l = new StringBuilder();
                    _l.Append(s);
                    _l.Append(",CustomAdRules");
                    ans.Add(_l.ToString());
                }
            }

            var outWriter = new StreamWriter(outFilePath);
            foreach (var i in ans)
            {
                outWriter.WriteLine(i);
            }

            in3Reader.Close();
            outWriter.Close();

            Console.WriteLine("完成");
        }
    }
}
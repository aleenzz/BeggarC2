using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace C2
{
    class Config
    {
        public static string Url = "http://192.168.0.103:8080";
        public static string InitUrl = "/AgentName/";
        public static string ShellUrl = "/AgentShell/";
        public static string PostResults = "/PostResults/";
        public static string key = "ac59075b964b0715";
        public static string GetAgentName()
        {
           return  Guid.NewGuid().ToString("N");
        }
        public static int Sleep = 6000;
    }
}

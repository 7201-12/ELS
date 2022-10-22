using Microsoft.AspNetCore.Mvc;
using ELS.Models;
using Redis.OM;
using Redis.OM.Searching;

namespace ELS.Controllers;


public class TestController : Controller{

    private readonly RedisCollection<Question> _question;
    private readonly RedisConnectionProvider _provider;
    public TestController(RedisConnectionProvider provider)
    {
        _provider = provider;
        _question = (RedisCollection<Question>)provider.RedisCollection<Question>();
    }

    [HttpPost]
    public Question AddQuestion([FromBody] Question question) {
        if (_question == null) {
            Console.WriteLine("NUll");
        }
        if (question == null) {
            Console.WriteLine("NUll q");
        }
        _question.Insert(question);
        return question;
    }

    public Q Question([FromBody] Q q) {
        Console.WriteLine(q.Id);
        Console.WriteLine(q.V);
        return q;
    }

    [HttpGet] 
    public IList<Question> GetQuestion([FromBody] string Id) {
        return _question.Where(x => x.Id == Id).ToList();
    }
}

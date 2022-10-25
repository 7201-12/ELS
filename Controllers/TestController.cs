using Microsoft.AspNetCore.Mvc;
using ELS.Models;
using System.Data;
using Npgsql;


namespace ELS.Controllers;


public class TestController : Controller{

    private readonly IConfiguration _configuration;
    public TestController(IConfiguration configuration) {
        _configuration = configuration;
    }

    [HttpGet]
    public JsonResult Get() {
        string query = @"
            select * from theory
        ";
        DataTable table = new DataTable();
        string strConn = _configuration.GetConnectionString("Els");
        Console.WriteLine(strConn);
        NpgsqlDataReader reader;
        using(NpgsqlConnection conn = new NpgsqlConnection(strConn)) {
            conn.Open();
            using(NpgsqlCommand command = new NpgsqlCommand(query, conn)) {
                reader = command.ExecuteReader();
                table.Load(reader);
                reader.Close();
                conn.Close();
            }
        }
        return new JsonResult(table);
    }

    [HttpPost]
    public Question AddQuestion() {
        Variant v1 = new Variant("1", "My name is Sir Lancelot of Camelot");
        Variant v2 = new Variant("1", "Sir Robin of Camelot");
        Variant v3 = new Variant("1", "It is 'Arthur', King of the Britons");
        List<Variant> l = new List<Variant>();
        l.Add(v1);
        l.Add(v2);
        l.Add(v3);
        Question q = new Question("1", "What is your name?", l, v3, 0.5);
        
        return q;
    }

    [HttpGet] 
    public Q Question([FromBody] Q q) {
        Console.WriteLine(q.Id);
        Console.WriteLine(q.V);
        return q;
    }
}

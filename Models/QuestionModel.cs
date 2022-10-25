namespace ELS.Models;


// public class Test {
//     [RedisIdField] [Indexed] public string? Id {get; set;}
//     [Indexed] public List<Question> Questions {get; set;}
// }

public class Question {
    public string? Id {get; set;}
    public string? Q {get; set;}
    public List<Variant> Variants {get; set;} = null!;
    public Variant Answer {get; set;} = null!;
    public double Time {get; set;}
    public Question(string id, string q, List<Variant> variants, Variant answer, double time) {
        Id = id;
        Q = q;
        Variants = variants;
        Answer = answer;
        Time = time;
    }
}

public class Variant {
    public string? QuestionId {get; set;}
    public string? Value {get; set;}
    public Variant(string id, string value) {
        QuestionId = id;
        Value = value;
    }
}

public class Q {
    public string? Id {get; set;}
    public string? V {get; set;}
}
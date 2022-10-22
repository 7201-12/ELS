using Redis.OM.Modeling;

namespace ELS.Models;


// public class Test {
//     [RedisIdField] [Indexed] public string? Id {get; set;}
//     [Indexed] public List<Question> Questions {get; set;}
// }

[Document(StorageType = StorageType.Json, Prefixes = new[]{"Question"})]
public class Question {
    [RedisIdField] [Indexed] public string? Id {get; set;}
    [Indexed] public string? Q {get; set;}
    [Indexed] public List<Variant> Variants {get; set;} = null!;
    [Indexed] public Variant Answer {get; set;} = null!;
}

public class Variant {
    [Indexed] public string? QuestionId {get; set;}
    [Indexed] public string? Value {get; set;}
}

public class Q {
    public string? Id {get; set;}
    public string? V {get; set;}
}
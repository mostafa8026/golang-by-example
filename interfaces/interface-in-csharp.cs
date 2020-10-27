public interface IShape{
    int x {get;set;}
    int y{get;set;}
    int Area();
}

public class Rectangle : IShape{
    public int x{get;set;}
    public int y {get;set;}
    public int Area(){
        return x*y;
    }
}

public class Circle : IShape{
    public int x{get;set;}
    public int y {get;set;}
    public int radious{get;set:}
    public int Area(){
        return Math.PI * Math.Pow(radious, 2);
    }
}

public int main(){
    IShape shape = new Circle();
    var c = shape as Circle;
    Console.WriteLine(c.radious);
    List<IShape> shapes = new List<IShape>();
    for (int i = 0; i<10; i++){
        if (i % 2 == 0){
            shapes.Add(new Circle());
        }else{
            shapes.Add(new Rectangle());
        }
    }

    foreach(var v in shapes){
        Console.WriteLine(v.Area())
    }
}
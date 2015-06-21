object FibonacciNumbers {

  def fibonacci(x:Int):Int = {
    if (x == 1) return 0

    if (x == 2) return 1

    return fibonacci(x - 1) + fibonacci(x - 2)
  }

  def main(args: Array[String]) = {
    println(fibonacci(readInt()))
  }

}
package base
import "github.com/ccxt/ccxt/go/v4"

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code


    func TestFetchLiquidations(exchange ccxt.ICoreExchange, skippedProperties interface{}, code interface{}) <- chan interface{} {
                ch := make(chan interface{})
                go func() interface{} {
                    defer close(ch)
                    defer ReturnPanicError(ch)
                        var method interface{} = "fetchLiquidations"
                if !IsTrue(GetValue(exchange.GetHas(), "fetchLiquidations")) {
            
                    ch <- true
                    return nil
                }
            
                items:= (<-exchange.FetchLiquidations(code))
                PanicOnError(items)
                Assert(IsArray(items), Add(Add(Add(Add(Add(Add(exchange.GetId(), " "), method), " "), code), " must return an array. "), exchange.Json(items)))
                // const now = exchange.Getmilliseconds() ();
                for i := 0; IsLessThan(i, GetArrayLength(items)); i++ {
                    TestLiquidation(exchange, skippedProperties, method, GetValue(items, i), code)
                }
                AssertTimestampOrder(exchange, method, code, items)
            
                ch <- true
                return nil
            
                }()
                return ch
            }

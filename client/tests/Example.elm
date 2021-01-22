module Example exposing (unitTest)

import Expect exposing (Expectation)
import Test exposing (..)


{-| See <https://github.com/elm-community/elm-test>
-}
unitTest : Test
unitTest =
    describe "simple unit test"
        [ test "1 + 1 = 2" <|
            \() ->
                1 + 1 |> Expect.equal 2
        ]

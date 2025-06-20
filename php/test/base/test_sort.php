<?php
namespace ccxt;

// ----------------------------------------------------------------------------

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code

// -----------------------------------------------------------------------------
include_once PATH_TO_CCXT . '/test/exchange/base/test_shared_methods.php';

function test_sort() {
    // todo: other argument checks
    $exchange = new \ccxt\Exchange(array(
        'id' => 'sampleexchange',
    ));
    $arr = ['b', 'a', 'c', 'd'];
    $sorted_arr = $exchange->sort($arr);
    assert_deep_equal($exchange, null, 'sort', $sorted_arr, ['a', 'b', 'c', 'd']);
}

each := func(seq, fn) {
    for x in seq { fn(x) }
}

export func(init, seq) {
    each(seq, func(x) { init += x })
    return init
}
fn remove_target() {
    let v1 = vec![1, 2, 2, 3, 4, 5, 6, 7, 2];

    let res: Vec<_> = v1.iter().filter(|&x| *x != 2).collect();

    for item in res {
        println!("{item}");
    } 
}

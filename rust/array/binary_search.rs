fn main() {
    let sorted_arr = [2, 4, 7, 8, 12];
    let target = 4;

    let mut low = 0;
    let mut hight = sorted_arr.len() - 1;

    while low <= hight {
        let mid = (low + hight) / 2;
        if target > sorted_arr[mid] {
            low = mid + 1;
        } else if target < sorted_arr[mid] {
            hight = mid - 1;
        } else {
            println!("find target");
            return;
        }
    }

    println!("not found");
}

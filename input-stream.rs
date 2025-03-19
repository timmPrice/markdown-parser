struct InputStream {
    input: String,
    pos: i32,
    line: i32,
    col: i32
}

impl Default for InputStream {
    fn default() -> Self {
        Self {
            input: String::new(),
            pos: 0,
            line: 1,
            col: 0,
        }
    }
}

fn main() {
    let input_stream = InputStream::default(); // Uses default values
    println!("Input: {}, Pos: {}, Line: {}, Col: {}", input_stream.input, input_stream.pos, input_stream.line, input_stream.col);
}

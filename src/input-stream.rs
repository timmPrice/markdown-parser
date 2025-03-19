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


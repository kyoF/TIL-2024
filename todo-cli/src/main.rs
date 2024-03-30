use std::collections::HashMap;

struct Todo {
    map: HashMap<String, bool>,
}

impl Todo {
    // // with text file
    // fn new() -> Result<Todo, std::io::Error> {
    //     let mut f = std::fs::OpenOptions::new()
    //         .write(true)
    //         .create(true)
    //         .read(true)
    //         .open("db.txt")?;
    //     let mut content = String::new();
    //     f.read_to_string(&mut content)?;
    //     let map: HashMap<String, bool> = content
    //         .lines()
    //         .map(|line| line.splitn(2, '\t').collect::<Vec<&str>>())
    //         .map(|v| (v[0], v[1]))
    //         .map(|(k, v)| (String::from(k), bool::from_str(v).unwrap()))
    //         .collect();
    //     Ok(Todo { map })
    // }
    // // use for loop with text file
    // fn new() -> Result<Todo, std::io::Error> {
    //     let mut f = std::fs::OpenOptions::new()
    //         .write(true)
    //         .create(true)
    //         .read(true)
    //         .open("db.txt")?;
    //     let mut content = String::new();
    //     f.read_to_string(&mut content)?;
    //     let mut map = HashMap::new();
    //     for entries in content.lines() {
    //         let mut values = entries.split('\t');
    //         let key = values.next().expect("No key");
    //         let val = values.next().expect("No value");
    //         map.insert(String::from(key), bool::from_str(val).unwrap());
    //     }
    // }
    fn new() -> Result<Todo, std::io::Error> {
        let f = std::fs::OpenOptions::new()
            .write(true)
            .create(true)
            .read(true)
            .open("db.json")?;
        match serde_json::from_reader(f) {
            Ok(map) => Ok(Todo { map }),
            Err(e) if e.is_eof() => Ok(Todo {
                map: HashMap::new(),
            }),
            Err(e) => panic!("An error occurred: {}", e),
        }
    }
    fn insert(&mut self, key: String) {
        self.map.insert(key, true);
    }
    // // with text file
    // fn save(self) -> Result<(), std::io::Error> {
    //     let mut content = String::new();
    //     for (k, v) in self.map {
    //         let record = format!("{}\t{}\n", k, v);
    //         content.push_str(&record)
    //     }
    //     std::fs::write("db.txt", content)
    // }
    // with json file
    fn save(self) -> Result<(), Box<dyn std::error::Error>> {
        let f = std::fs::OpenOptions::new()
            .write(true)
            .create(true)
            .open("db.json")?;
        serde_json::to_writer_pretty(f, &self.map)?;
        Ok(())
    }
    fn complete(&mut self, key: &String) -> Option<()> {
        match self.map.get_mut(key) {
            Some(v) => Some(*v = false),
            None => None,
        }
    }
}

fn main() {
    let action = std::env::args().nth(1).expect("Please specify an action");
    let item = std::env::args().nth(2).expect("Please specify an item");

    let mut todo = Todo::new().expect("Initialisation of todo failed");

    if action == "add" {
        todo.insert(item);
        match todo.save() {
            Ok(_) => println!("todo saved"),
            Err(why) => println!("An error occurred: {}", why),
        }
    } else if action == "complete" {
        match todo.complete(&item) {
            None => println!("'{}' is not present in the list", item),
            Some(_) => match todo.save() {
                Ok(_) => println!("todo saved"),
                Err(why) => println!("An error occurred: {}", why),
            },
        }
    }
}

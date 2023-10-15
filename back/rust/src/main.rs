use postgres::{Client, NoTls};
use postgres::Error as PostgresError;
use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};
use std::collections::HashMap;


#[macro_use]
extern crate serde_derive;

#[derive(Serialize, Deserialize)]
struct Produto {
    id: Option<i32>,
    nome: String,
    descricao: String,
    preco: f64,
    vendido: bool,
}

//DATABASE_URL
const DB_URL: &str = "postgres://postgres:postgres@pgbouncer:6432/postgres";

//constants
const OK_RESPONSE: &str = "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n";
const NOT_FOUND: &str = "HTTP/1.1 404 NOT FOUND\r\n\r\n";
const INTERNAL_SERVER_ERROR: &str = "HTTP/1.1 500 INTERNAL SERVER ERROR\r\n\r\n";

//main function
fn main() {
    //Set database
    if let Err(e) = set_database() {
        println!("Error: {}", e);
        return;
    }


    //start server and print port
    let listener = TcpListener::bind(format!("0.0.0.0:8060")).unwrap();
    println!("Server started at port 8060");

    //handle the client
    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                handle_client(stream);
            }
            Err(e) => {
                println!("Error: {}", e);
            }
        }
    }
}

//handle_client function
fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    let mut request = String::new();

    match stream.read(&mut buffer) {
        Ok(size) => {
            request.push_str(String::from_utf8_lossy(&buffer[..size]).as_ref());

            let (status_line, mut content) = match &*request {
                r if r.starts_with("POST /produtos") => handle_post_request(r),
                r if r.starts_with("GET /produtos/") => handle_get_request(r),
                r if r.starts_with("GET /produtos") => handle_get_all_request(r),
                r if r.starts_with("PUT /produtos/") => handle_put_request(r),
                r if r.starts_with("PATCH /produtos/") => handle_patch_request(r),
                r if r.starts_with("DELETE /produtos/") => handle_delete_request(r),
                _ => (NOT_FOUND.to_string(), "404 Not Found".to_string()),
            };

            // Add the header to the response
            content = format!("{}Backend-Type: Rust\r\n\r\n{}", content, "");

            stream.write_all(format!("{}{}", status_line, content).as_bytes()).unwrap();

        }
        Err(e) => {
            println!("Error: {}", e);
        }
    }
}

//CONTROLLERS

//handle_post_request function
fn handle_post_request(request: &str) -> (String, String) {
    match (get_user_request_body(&request), Client::connect(DB_URL, NoTls)) {
        (Ok(user), Ok(mut client)) => {
            client
                .execute(
                    "INSERT INTO produtos (nome
                , descricao, preco) VALUES ($1, $2,$3)",
                    &[&user.nome
            , &user.descricao, &user.preco],
                )
                .unwrap();

            (OK_RESPONSE.to_string(), "Product created".to_string())
        }
        _ => (INTERNAL_SERVER_ERROR.to_string(), "Error".to_string()),
    }
}

//handle_get_request function
fn handle_get_request(request: &str) -> (String, String) {
    match (get_id(&request).parse::<i32>(), Client::connect(DB_URL, NoTls)) {
        (Ok(id), Ok(mut client)) =>
            match client.query_one("SELECT * FROM produtos WHERE id = $1", &[&id]) {
                Ok(row) => {
                    let user = Produto {
                        id: row.get(0),
                        nome
                : row.get(1),
                        descricao: row.get(2),
                        preco: row.get(3),
                        vendido: row.get(4),
                    };

                    (OK_RESPONSE.to_string(), serde_json::to_string(&user).unwrap())
                }
                _ => (NOT_FOUND.to_string(), "Produto não encontrado".to_string()),
            }

        _ => (INTERNAL_SERVER_ERROR.to_string(), "Error".to_string()),
    }
}

//handle_get_all_request function
fn handle_get_all_request(_request: &str) -> (String, String) {
    match Client::connect(DB_URL, NoTls) {
        Ok(mut client) => {
            let mut users = Vec::new();

            for row in client.query("SELECT * FROM produtos", &[]).unwrap() {
                users.push(Produto {
                    id: row.get(0),
                    nome
            : row.get(1),
                    descricao: row.get(2),
                    preco: row.get(3),
                    vendido: row.get(4),
                });
            }

            (OK_RESPONSE.to_string(), serde_json::to_string(&users).unwrap())
        }
        _ => (INTERNAL_SERVER_ERROR.to_string(), "Error Not Found".to_string()),
    }
}

//handle_put_request function
fn handle_put_request(request: &str) -> (String, String) {
    match
    (
        get_id(&request).parse::<i32>(),
        get_user_request_body(&request),
        Client::connect(DB_URL, NoTls),
    )
    {
        (Ok(id), Ok(user), Ok(mut client)) => {
            client
                .execute(
                    "UPDATE produtos SET nome
             = $1, descricao = $2, preco = $3 WHERE id = $4",
                    &[&user.nome
            , &user.descricao, &user.preco, &id],
                )
                .unwrap();

            (OK_RESPONSE.to_string(), "User updated".to_string())
        }
        _ => (INTERNAL_SERVER_ERROR.to_string(), "Error".to_string()),
    }
}

//handle_patch_request function
fn handle_patch_request(request: &str) -> (String, String) {
    match
    (
        get_id(&request).parse::<i32>(),
        get_user_request_body(&request),
        Client::connect(DB_URL, NoTls),
    )
    {
        (Ok(id), Ok(user), Ok(mut client)) => {
            client
                .execute(
                    "UPDATE produtos SET vendido = $1 WHERE id = $2",
                    &[&user.vendido, &id],
                )
                .unwrap();

            (OK_RESPONSE.to_string(), "User updated".to_string())
        }
        _ => (INTERNAL_SERVER_ERROR.to_string(), "Error".to_string()),
    }
}

//handle_delete_request function
fn handle_delete_request(request: &str) -> (String, String) {
    match (get_id(&request).parse::<i32>(), Client::connect(DB_URL, NoTls)) {
        (Ok(id), Ok(mut client)) => {
            let rows_affected = client.execute("DELETE FROM produtos WHERE id = $1", &[&id]).unwrap();

            if rows_affected == 0 {
                return (NOT_FOUND.to_string(), "User not found".to_string());
            }

            (OK_RESPONSE.to_string(), "User deleted".to_string())
        }
        _ => (INTERNAL_SERVER_ERROR.to_string(), "Error".to_string()),
    }
}

//set_database function
fn set_database() -> Result<(), PostgresError> {
    //Connect to database
    let mut client = Client::connect(DB_URL, NoTls)?;

    //Create table
    client.batch_execute(
        "CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            nome
     VARCHAR NOT NULL,
            email VARCHAR NOT NULL
        )"
    )?;
    Ok(())
}

//get_id function
fn get_id(request: &str) -> &str {
    request.split("/").nth(2).unwrap_or_default().split_whitespace().next().unwrap_or_default()
}

//deserialize user from request body with the id
fn get_user_request_body(request: &str) -> Result<Produto, serde_json::Error> {
    serde_json::from_str(request.split("\r\n\r\n").last().unwrap_or_default())
}
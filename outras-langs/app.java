import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;

public class ConsumirAPIREST {

    public static void main(String[] args) {
        // Listar todos os produtos
        listarProdutos();

        // Obter detalhes de um produto por ID
        obterDetalhesDoProduto(1);

        // Criar um novo produto
        criarProduto("Novo Produto", "Descrição do Novo Produto", 39.99);

        // Atualizar um produto existente (Substituir o ID pelo ID do produto real)
        atualizarProduto(1, "Produto Atualizado", "Descrição do Produto Atualizado", 49.99);

        // Excluir um produto por ID (Substituir o ID pelo ID do produto real)
        excluirProduto(2);
    }

    public static void listarProdutos() {
        try {
            URL url = new URL("http://localhost:5000/produtos");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");
            conn.setRequestProperty("Accept", "application/json");

            if (conn.getResponseCode() != 200) {
                throw new RuntimeException("Erro HTTP: " + conn.getResponseCode());
            }

            BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            String output;
            System.out.println("Lista de Produtos:");
            while ((output = br.readLine()) != null) {
                System.out.println(output);
            }

            conn.disconnect();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void obterDetalhesDoProduto(int id) {
        try {
            URL url = new URL("http://localhost:5000/produtos/" + id);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("GET");
            conn.setRequestProperty("Accept", "application/json");

            if (conn.getResponseCode() != 200) {
                throw new RuntimeException("Erro HTTP: " + conn.getResponseCode());
            }

            BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            String output;
            System.out.println("Detalhes do Produto:");
            while ((output = br.readLine()) != null) {
                System.out.println(output);
            }

            conn.disconnect();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void criarProduto(String nome, String descricao, double preco) {
        try {
            URL url = new URL("http://localhost:5000/produtos");
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("POST");
            conn.setRequestProperty("Content-Type", "application/json");
            conn.setDoOutput(true);

            String input = String.format("{\"nome\":\"%s\",\"descricao\":\"%s\",\"preco\":%.2f}", nome, descricao, preco);

            OutputStream os = conn.getOutputStream();
            os.write(input.getBytes());
            os.flush();

            if (conn.getResponseCode() != 201) {
                throw new RuntimeException("Erro HTTP: " + conn.getResponseCode());
            }

            System.out.println("Novo Produto criado com sucesso!");

            conn.disconnect();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void atualizarProduto(int id, String nome, String descricao, double preco) {
        try {
            URL url = new URL("http://localhost:5000/produtos/" + id);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("PUT");
            conn.setRequestProperty("Content-Type", "application/json");
            conn.setDoOutput(true);

            String input = String.format("{\"nome\":\"%s\",\"descricao\":\"%s\",\"preco\":%.2f}", nome, descricao, preco);

            OutputStream os = conn.getOutputStream();
            os.write(input.getBytes());
            os.flush();

            if (conn.getResponseCode() != 200) {
                throw new RuntimeException("Erro HTTP: " + conn.getResponseCode());
            }

            System.out.println("Produto Atualizado com sucesso!");

            conn.disconnect();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void excluirProduto(int id) {
        try {
            URL url = new URL("http://localhost:5000/produtos/" + id);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("DELETE");

            if (conn.getResponseCode() != 200) {
                throw new RuntimeException("Erro HTTP: " + conn.getResponseCode());
            }

            System.out.println("Produto Excluído com sucesso!");

            conn.disconnect();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

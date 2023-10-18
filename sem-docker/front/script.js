// Endereço do back
const url = 'http://localhost:5000';

// Função para mostrar Backend-Type
function mostrarBackendType(response) {
    const header = document.getElementById('header');
    header.innerHTML = `Backend: ${response.headers.get('Backend-Type')}`;
    console.log(response.headers.get('Backend-Type'));
    return response.json()
}

// Função para listar todos os produtos
function listarProdutos() {
    fetch(url + '/produtos')
        .then(response => mostrarBackendType(response))
        .then(data => {
            const productList = document.getElementById('product-list');
            productList.innerHTML = '';
            data.forEach(product => {
                const listItem = document.createElement('li');
                listItem.innerHTML = `<a href="#" onclick="mostrarDetalhes(${product.id})">${product.nome}</a>`;
                productList.appendChild(listItem);
            });
        });
}

// Função para mostrar detalhes de um produto por ID
function mostrarDetalhes(id) {
    fetch(url + `/produtos/${id}`)
        .then(response => mostrarBackendType(response))
        .then(product => {
            const productDetails = document.getElementById('product-details');
            productDetails.innerHTML = `ID: ${product.id}<br>Nome: ${product.nome}<br>Descrição: ${product.descricao}<br>Preço: ${product.preco}`;
        });
}

// Função para criar um novo produto
document.getElementById('create-form').addEventListener('submit', function (event) {
    event.preventDefault();
    const name = document.getElementById('create-name').value;
    const description = document.getElementById('create-description').value;
    const price = parseFloat(document.getElementById('create-price').value);

    fetch(url + '/produtos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nome: name, descricao: description, preco: price })
    }).then(response => mostrarBackendType(response))
        .then(() => {
            listarProdutos();
        });
});

// Função para atualizar um produto existente
document.getElementById('update-form').addEventListener('submit', function (event) {
    event.preventDefault();
    const id = parseInt(document.getElementById('update-id').value);
    const name = document.getElementById('update-name').value;
    const description = document.getElementById('update-description').value;
    const price = parseFloat(document.getElementById('update-price').value);

    fetch(url + `/produtos/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nome: name, descricao: description, preco: price })
    }).then(response => mostrarBackendType(response))
        .then(() => {
            listarProdutos();
        });
});

// Função para excluir um produto por ID
document.getElementById('delete-form').addEventListener('submit', function (event) {
    event.preventDefault();
    const id = parseInt(document.getElementById('delete-id').value);

    fetch(url + `/produtos/${id}`, {
        method: 'DELETE'
    }).then(response => mostrarBackendType(response))
        .then(() => {
            listarProdutos();
        });
});

// Listar produtos ao carregar a página
listarProdutos();
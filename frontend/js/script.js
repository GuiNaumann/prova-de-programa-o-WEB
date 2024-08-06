
// quando clicar submit no html ele manda uma ação para o componente forme q tem o id cadastro-form
//  submit poderia ser outras coisas como hover clique etc
// async chamo por conta q estou chamando uma api
// toda função q tenho q agarda o resultado da api tem q ser async
document.getElementById('cadastro-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    const nome = document.getElementById('nome').value;
    const telefone = document.getElementById('telefone').value;

    // tratamento de erro simples
    //  se der erro no try ele cai no catch
    try {
        // requisita a API apenas
        //  await espera a resposta da api
        const response = await fetch('http://localhost:8080/aluno', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            // trasformei o objeto em string json
            // pq o corpo de uma api REST tem q ser json
            body: JSON.stringify({ nome, telefone })
        });

        const aluno = await response.json();
        alert('Aluno cadastrado com sucesso!');
        // depois que cadastra faz o get da lista completa de alunos
        await fetchAlunos();
    } catch (error) {
        console.error('Erro ao cadastrar aluno:', error);
        alert('Erro ao cadastrar aluno.');
    }
});

async function fetchAlunos() {
    try {
        // respomse é o get da API
        const response = await fetch('http://localhost:8080/aluno');
        const alunos = await response.json();
         // pega o componente UL que tem o ID aluno-list
         const alunoList = document.getElementById('aluno-list');
        alunoList.innerHTML = '';
        alunos.forEach(aluno => {
            const listItem = document.createElement('li');
            listItem.textContent = `ID: ${aluno.ID}, Nome: ${aluno.Nome}, Telefone: ${aluno.Telefone}`;
            alunoList.appendChild(listItem);
        });
    } catch (error) {
        console.error('Erro ao buscar alunos:', error);
    }
}

fetchAlunos();

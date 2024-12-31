<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Model Manager</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #282c34;
            color: #abb2bf;
            margin: 0;
            padding: 0;
        }
        h1 {
            background-color: #61afef;
            color: #282c34;
            padding: 10px;
            text-align: center;
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 20px;
        }
        th, td {
            border: 1px solid #abb2bf;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #61afef;
            color: #282c34;
        }
        tr:nth-child(even) {
            background-color: #3e4451;
        }
        tr:hover {
            background-color: #5c6370;
        }
        .button {
            background-color: #61afef;
            color: #282c34;
            border: none;
            padding: 10px 20px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 4px 2px;
            cursor: pointer;
            border-radius: 4px;
        }
        .button:hover {
            background-color: #98c379;
        }
        .form-container {
            margin-top: 20px;
            padding: 20px;
            background-color: #3e4451;
            border-radius: 4px;
        }
        .form-field {
            margin-bottom: 10px;
        }
        .form-field label {
            display: block;
            margin-bottom: 5px;
        }
        .form-field input, .form-field select {
            width: 100%;
            padding: 8px;
            border: 1px solid #abb2bf;
            border-radius: 4px;
            background-color: #282c34;
            color: #abb2bf;
        }
        #messages {
            margin-top: 20px;
            color: #98c379;
        }
    </style>
</head>
<body>
    <h1>Model Manager</h1>
    <button class="button" onclick="loadModels()">Load Models</button>

    <div class="form-container">
        <h2>Create New Model</h2>
        <div class="form-field">
            <label for="name">Name:</label>
            <input type="text" id="name" name="name">
        </div>
        <div class="form-field">
            <label for="path">Path:</label>
            <input type="text" id="path" name="path">
        </div>
        <div class="form-field">
            <label for="size">Size:</label>
            <input type="number" id="size" name="size">
        </div>
        <div class="form-field">
            <label for="param">Param:</label>
            <input type="number" id="param" name="param">
        </div>
        <div class="form-field">
            <label for="quant">Quant:</label>
            <input type="text" id="quant" name="quant">
        </div>
        <div class="form-field">
            <label for="gpuLayers">GPU Layers:</label>
            <input type="number" id="gpuLayers" name="gpuLayers">
        </div>
        <div class="form-field">
            <label for="gpuLayersUsed">GPU Layers Used:</label>
            <input type="number" id="gpuLayersUsed" name="gpuLayersUsed">
        </div>
        <div class="form-field">
            <label for="ctx">Ctx:</label>
            <input type="number" id="ctx" name="ctx">
        </div>
        <div class="form-field">
            <label for="ctxUsed">Ctx Used:</label>
            <input type="number" id="ctxUsed" name="ctxUsed">
        </div>
        <button class="button" onclick="createModel()">Create Model</button>
    </div>

    <div id="messages"></div>

    <table id="modelsTable">
        <thead>
            <tr>
                <th>Name</th>
                <th>Path</th>
                <th>Size</th>
                <th>Param</th>
                <th>Quant</th>
                <th>GPU Layers</th>
                <th>GPU Layers Used</th>
                <th>Ctx</th>
                <th>Ctx Used</th>
                <th>King Of</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            <!-- Models will be loaded here -->
        </tbody>
    </table>

    <script>
        function loadModels() {
            fetch('/models')
                .then(response => response.json())
                .then(models => {
                    const tableBody = document.getElementById('modelsTable').getElementsByTagName('tbody')[0];
                    tableBody.innerHTML = '';
                    models.forEach(model => {
                        const row = tableBody.insertRow();
                        row.insertCell().textContent = model.name;
                        row.insertCell().textContent = model.path;
                        row.insertCell().textContent = model.size;
                        row.insertCell().textContent = model.param;
                        row.insertCell().textContent = model.quant;
                        row.insertCell().textContent = model.gpuLayers;
                        row.insertCell().textContent = model.gpuLayersUsed;
                        row.insertCell().textContent = model.ctx;
                        row.insertCell().textContent = model.ctxUsed;
                        row.insertCell().textContent = model.kingOf;
                        const actionsCell = row.insertCell();
                        const editButton = document.createElement('button');
                        editButton.textContent = 'Edit';
                        editButton.className = 'button';
                        editButton.onclick = () => editModel(model.name);
                        actionsCell.appendChild(editButton);
                        const deleteButton = document.createElement('button');
                        deleteButton.textContent = 'Delete';
                        deleteButton.className = 'button';
                        deleteButton.onclick = () => deleteModel(model.name);
                        actionsCell.appendChild(deleteButton);
                    });
                })
                .catch(error => {
                    console.error('Error:', error);
                    document.getElementById('messages').textContent = 'Failed to load models.';
                });
        }

        function createModel() {
            const model = {
                name: document.getElementById('name').value,
                path: document.getElementById('path').value,
                size: parseFloat(document.getElementById('size').value),
                param: parseFloat(document.getElementById('param').value),
                quant: document.getElementById('quant').value,
                gpuLayers: parseInt(document.getElementById('gpuLayers').value),
                gpuLayersUsed: parseInt(document.getElementById('gpuLayersUsed').value),
                ctx: parseInt(document.getElementById('ctx').value),
                ctxUsed: parseInt(document.getElementById('ctxUsed').value)
            };

            fetch('/models', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(model)
            })
            .then(response => {
                if (response.ok) {
                    document.getElementById('messages').textContent = 'Model created successfully.';
                    loadModels();
                } else {
                    throw new Error('Failed to create model.');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                document.getElementById('messages').textContent = 'Failed to create model.';
            });
        }

        function editModel(name) {
            // Implement edit functionality here
            console.log('Edit model:', name);
        }

        function deleteModel(name) {
            if (!confirm(`Are you sure you want to delete the model "${name}"?`)) {
                return;
            }
            fetch(`/models/${name}`, {
                method: 'DELETE'
            })
            .then(response => {
                if (response.ok) {
                    document.getElementById('messages').textContent = 'Model deleted successfully.';
                    loadModels();
                } else {
                    throw new Error('Failed to delete model.');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                document.getElementById('messages').textContent = 'Failed to delete model.';
            });
        }
    </script>
</body>
</html>

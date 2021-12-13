const confirmDelete = {
    setup: () => {
        confirmDelete.bindDeleteItem();
        confirmDelete.hideModal();
    },
    modal: document.getElementById('deleteModal'),
    deleteItemForm: document.getElementById('deleteItem'),
    modalBodyText: document.getElementById('modalBodyText'),
    bindDeleteItem: () => {
        document.querySelectorAll('.delete-item').forEach(item => {
            item.addEventListener('click', () => {
                let formAction = item.getAttribute('data-url');
                let entity = item.getAttribute('data-entity');
                let textModalBody = 'Are you sure you want to delete this ' + entity + '?';
                confirmDelete.deleteItemForm.setAttribute('action', formAction);
                confirmDelete.modalBodyText.innerHTML = textModalBody;
            });
        });
    },
    hideModal: () => {
        $(confirmDelete.modal).on('hidden.bs.modal', () => {
            confirmDelete.deleteItemForm.setAttribute('action', '');
        })
    }
}

export default confirmDelete;

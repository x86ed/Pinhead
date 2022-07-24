import { getUsers } from "./api.js";
import { confirmDeleteUser } from "./admin.js";

const TABLE_SORT_DIRECTION = {
  NONE: 'none',
  ASCENDING: 'ascending',
  DESCENDING: 'descending',
};
 
const columns = [
  {
    id: 'name',
    title: 'Name',
    sortCycle: 'tri-states-from-ascending',
  },
  {
    id: 'role',
    title: 'Role',
    sortCycle: 'tri-states-from-ascending',
  },
];

var rows;

// event listener for dynamic delete button created per row as needed
document.querySelector('body').addEventListener('click', function (e) {
  var userId = null;
  var userName = null;
  if (e.target.classList.contains('confirmDeleteUserBtn')) {
    var btn = e.target;
    userName = btn.getAttribute('username');
    userId = btn.getAttribute('userId');
  }
  if (e.target.parentNode.classList.contains('confirmDeleteUserBtn')) {
    var btn = e.target.parentNode;
    userName = btn.getAttribute('username');
    userId = btn.getAttribute('userId');
  }
  
  if (userName && userId) {
    confirmDeleteUser(userName, userId);
  }
});

// event listener for the table sorting event
document.addEventListener('bx-table-header-cell-sort', ({ defaultPrevented, detail, target }) => {
  if (!defaultPrevented) {
    const { columnId } = target.dataset;
    const { sortDirection: direction } = detail;
    // Sets the sorting as user desires
    const sortInfo = {
      columnId,
      direction,
    };
  
    table.state.setSortInfo(sortInfo);
  }
});

// event listener for pagination event
document.addEventListener('bx-pagination-changed-current', ({ detail }) => {
  table.state.setStart(detail.start);
});
  
// returns pagination component
function _renderPagination() {
  const { pageSize, start } = table.state;
  if (typeof pageSize === 'undefined') {
    return undefined;
  }
  
  return `
    <bx-pagination
      class="paginationNoScroll"
      page-size="${pageSize}"
      start="${start}"
      total="${rows.length}"
    >
    </bx-pagination>
  `;
}
  
const collator = new Intl.Collator('en');
  
export const table = () => {
  const { sortInfo, start, pageSize } = table.state;
  const { columnId: sortColumnId, direction: sortDirection } = sortInfo;
  // sorting logic, returns the sorted rows to render
  const sortedRows = rows.slice().sort((lhs, rhs) => {
  const lhsValue = lhs[sortInfo.columnId];
  const rhsValue = rhs[sortInfo.columnId];
    return (sortInfo.direction === 'ascending' ? 1 : -1) * collator.compare(lhsValue, rhsValue);
  });
  
  if (typeof pageSize === 'undefined') {
    return undefined;
  }
  
  return `
    <bx-table size="compact">
      <bx-table-head>
        <bx-table-header-row>
          ${columns
            .map(column => {
              const { id: columnId, sortCycle, title } = column;
              const sortDirectionForThisCell =
                sortCycle && (columnId === sortColumnId ? sortDirection : TABLE_SORT_DIRECTION.NONE);
              return `<bx-table-header-cell ${sortCycle && `sort-cycle="${sortCycle}"`} ${
                sortDirectionForThisCell && `sort-direction="${sortDirectionForThisCell}"`
              } data-column-id="${columnId}">${title}</bx-table-header-cell>`;
            })
            .join('')}
          <bx-table-header-cell></<bx-table-header-cell>
        </bx-table-header-row>
      </bx-table-head>
      <bx-table-body id="tableBody">
        ${sortedRows
          .slice(start, start + pageSize)
          .map(row => {
            const { id: rowId } = row;
            return `
        <bx-table-row data-row-id="${rowId}">
          ${columns
            .map(column => {
              const { id: columnId } = column;
                return `<bx-table-cell>${row[columnId]}</bx-table-cell>`;
              })
            .join('')}
          <bx-table-cell>
            <button class="confirmDeleteUserBtn deleteBtn" username="${row.name}" userId="${row.id}">
              <i class="fa fa-trash"></i>
            </button>
          </bx-table-cell>
        </bx-table-row>`;
        })
        .join('')}
      </bx-table-body>
    </bx-table>
  ${_renderPagination()}
  `;
};
  
table.state = {
  start: 0,
  setStart: start => {
    setState(() => {
      table.state.start = start;
    });
  },
  pageSize: 10,
  setPageSize: pageSize => {
    setState(() => {
      table.state.pageSize = pageSize;
    });
  },
  sortInfo: {
    columnId: 'name',
    direction: TABLE_SORT_DIRECTION.ASCENDING,
  },
  setSortInfo: sortInfo => {
    setState(() => {
      table.state.sortInfo = sortInfo;
    });
  },
};
  
  function setState(callback) {
    callback();
    updateTree(); // extracted function
  }

async function updateTree() {
  var tableReady = document.getElementById('users-table');
  if (tableReady) {
    var users = await getUsers();
    if (users == null || users == undefined || users.is_error) {
      console.log("token has expired");
      window.location.href="signin.html";
    }
    rows = users.map(function (user) {
      return {
        id: user.ID,
        name: user.name,
        role: user.role
      }
    });
    tableReady.innerHTML = table();
  }
};

await updateTree();
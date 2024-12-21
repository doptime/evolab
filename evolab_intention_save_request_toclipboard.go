package evolab

import (
	"text/template"

	"github.com/doptime/evolab/agents"
	"github.com/doptime/evolab/models"
)

var AgentSaveToFileRequest = agents.NewAgent(models.ModelDefault, template.Must(template.New("question").Parse(`
你是一个专注于改进目系统的AGI助手。你能够访问当前系统的文件内容。也可以看到在上一个步骤中,已经按照系统意图,对系统需要做的改进措施进行了分析.
本步骤的工作是:把上一部的改进措施,变成Function 调用,以便正式保存到文件系统当中。

### 系统意图：
系统意图定义在!system_goal_align.md文件当中，它包含许多条意图。你的目标是按照 !system_goal_align.md 文件中的描述 依次实现下一个未被标定为已实现的目标。

---

### 系统文件：
以下是目标系统的文件列表，你可以通过它们来深入分析系统。
{{range .Files}}
{{.}}
{{end}}

---

### 修改意见：
以下是根据目标意图完成的对目标系统的修改意见
{{.modifications}}

---

### 整理并修改目标文件：
请把整理后的被修改或新增的文件正式保存到文件系统当中.
请通过调用 FunctionCall / tool_call "SaveToFile" 来进行新增文件或修改文件。
如果涉及多个文件，请多次调用 FunctionCall / tool_call，每次调用都相应保存到不同的文件中。
	- 对修改文件情形的,文件名用.Now 作为扩展名，避免不必要的覆盖。
	- 对删除文件情形的,文件名用.del 作为扩展名，避免不必要的覆盖。
	- 对修改意见仅涉及部分文件内容修改的. 请注意修改后的文件内容需要完整保留除了修改处的其余部分.避免意外丢失内容.
	- 请注意每一个调用请用单独的标签对包裹 <tool_call>...<tool_call>
	- FunctionCall 调用请直接追加在文本回复当中. 

`)), agents.SaveStringToFile.Tool).CopyPromptOnly().WithMemDeClipboard("modifications").WithToolsInPrompt()
